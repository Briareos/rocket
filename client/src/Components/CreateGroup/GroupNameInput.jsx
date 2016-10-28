import React, {Component, PropTypes} from 'react';

import GroupList from './GroupList';
import GroupItem from './GroupItem';
import CreateGroupForm from './CreateGroupForm';

class GroupNameInput extends Component {
    constructor(props) {
        super(props);

        this.state = {
            existingGroups: [],
            showCreateForm: false,
            currentlySelected: null,
            query: ''
        };

        this.handleInputUpdate = this.handleInputUpdate.bind(this);
        this.handleKeyPress = this.handleKeyPress.bind(this);
        this.handleGroupSelect = this.handleGroupSelect.bind(this);
        this.handleFormSubmit = this.handleFormSubmit.bind(this);
    }
    render() {
        return (
            <div className="group-name-input-wrapper">
                <input type="text" onChange={this.handleInputUpdate} onKeyDown={this.handleKeyPress}/>
                {this.state.existingGroups.length > 0 && !this.state.showCreateForm ?
                    <GroupList>
                        {this.state.existingGroups.map(group =>
                            <GroupItem key={group.id} group={group} onClick={this.handleGroupSelect.bind(this, group)} selected={group.id === this.state.currentlySelected}/>
                        )}
                    </GroupList>
                : this.state.query.length > 0 ?
                    <CreateGroupForm onFormSubmit={this.handleFormSubmit}/>
                : ''}
            </div>
        )
    }
    handleFormSubmit(groupForm) {
        var newGroup = Object.assign({}, groupForm, {name: this.state.query});
        this.props.onEnter(
            {
                type: 'create',
                group: newGroup
            }
        );
    }
    handleInputUpdate(event) {
        let query = event.target.value;
        let existingGroups = [];

        if (query) {
            existingGroups = this.props.groups.filter(group => group.name.includes(query));
        }

        this.setState(Object.assign({}, this.state, {
            showCreateForm: false,
            existingGroups,
            query
        }));
    }
    handleKeyPress(event) {
        if(event.key == 'Enter') {
            var existingGroup = this.props.groups.filter(group => group.name === this.state.query);

            if (existingGroup.length === 0) {
                this.setState(Object.assign({}, this.state, {showCreateForm: true}));
            } else {
                this.props.onEnter(
                    {
                        type: 'select',
                        group: existingGroup[0]
                    }
                );
            }
        }
    }
    handleGroupSelect(group) {
        this.props.onEnter(
            {
                type: 'select',
                group
            }
        );
    }
}

GroupNameInput.propTypes = {
    groups: PropTypes.array.isRequired,
    onEnter: PropTypes.func
};

export default GroupNameInput;