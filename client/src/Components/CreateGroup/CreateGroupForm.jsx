import React, {Component, PropTypes} from 'react';

class CreateGroupForm extends Component {
    constructor(props) {
        super(props);

        this.handleCreate = this.handleCreate.bind(this);
    }
    render() {
        return (
            <div className="create-group-form">
                <textarea ref="groupDescription"></textarea>
                <input type="checkbox" ref="busyValue"/>
                <input type="checkbox" ref="remoteValue"/>
                <button onClick={this.handleCreate}>Create Group</button>
            </div>
        );
    }
    handleCreate() {
        let group = {
            description: this.refs.groupDescription.value,
            busyValue: this.refs.busyValue.checked,
            remoteValue: this.refs.remoteValue.checked
        };
        this.props.onFormSubmit(group);
    }
}

CreateGroupForm.propTypes = {
    onFormSubmit: PropTypes.func.isRequired
};

export default CreateGroupForm;