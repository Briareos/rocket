import React, {Component} from 'react';
import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';
import { browserHistory } from 'react-router';

import * as Actions from '../../actions/actionCreators';
import GroupNameInput from './GroupNameInput';

class CreateGroupPage extends Component {
    constructor(props) {
        super(props);

        this.handleInputEnter = this.handleInputEnter.bind(this);
    }
    render() {

        return (
            <div id="group-page" className="page">
                <GroupNameInput groups={this.props.groups} onEnter={this.handleInputEnter}/>
            </div>
        );
    }
    handleInputEnter(action) {
        console.log(action);
        switch (action.type) {
            case 'select':
                browserHistory.push('/group/' + action.group.id);
                break;
            case 'create':
                let group = action.group;
                this.props.actions.createGroup(
                    group.name,
                    group.description,
                    group.busyValue,
                    group.remoteValue
                );
                break;
        }
    }
}

function mapStateToProps (state) {
    return {
        groups: state.groups
    };
}

function mapDispatchToProps(dispatch) {
    return {
        actions: bindActionCreators(Actions, dispatch)
    }
}
export default connect(mapStateToProps, mapDispatchToProps)(CreateGroupPage);