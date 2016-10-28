import React, {Component} from 'react';
import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';

import * as Actions from '../../actions/actionCreators';

class GroupPage extends Component {
    constructor(props) {
        super(props);
    }
    render() {
        return (
            <div id="group-page" className="page">
                <GroupNameInput groups={this.props.groups}/>
            </div>
        );
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
export default connect(mapStateToProps, mapDispatchToProps())(GroupPage);