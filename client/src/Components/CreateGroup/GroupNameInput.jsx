import React, {Component, PropTypes} from 'react';

import GroupList from './GroupList';
import GroupItem from './GroupItem';

class GroupNameInput extends Component {
    constructor(props) {
        super(props);

        this.state = {
            existingGroups: []
        };

        this.handleInputUpdate = this.handleInputUpdate.bind(this);
        this.handleKeyPress = this.handleKeyPress.bind(this);
        this.handleGroupSelect = this.handleGroupSelect.bind(this);
    }
    render() {
        return (
            <div className="group-name-input-wrapper">
                <input type="text" onChange={this.handleInputUpdate} onKeyDown={this.handleKeyPress}/>
                {this.state.existingGroups.length > 0 ?
                    <GroupList>
                        {this.state.existingGroups.map(group =>
                            <GroupItem group={group} onSelect={this.handleGroupSelect}/>
                        )}
                    </GroupList>
                : ''}
            </div>
        )
    }
    handleInputUpdate(event) {
        var query = event.target.value;

        if (query) {
            this.setState({
                existingGroups: this.props.groups.filter(group => group.name.includes(query))
            });
        } else {
            this.setState({
                existingGroups: []
            });
        }
    }
    handleKeyPress(event) {
        if(event.key == 'Enter') {
            console.log('enter press here! ')
        }
    }
    handleGroupSelect() {

    }
}

GroupNameInput.propTypes = {
    groups: PropTypes.array.isRequired,
    onEnter: PropTypes.func
};

export default GroupNameInput;