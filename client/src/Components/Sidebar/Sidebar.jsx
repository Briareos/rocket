import React, {Component} from 'react';
import {connect} from 'react-redux';

import SidebarUserInfo from './SidebarUserInfo';

class Sidebar extends Component {
    constructor(props) {
        super(props);
    }
    render() {
        return (
            <div className="sidebar">
                <SidebarUserInfo user={this.props.user} />
            </div>
        );
    }
}

function mapStateToProps(state) {
    return {
        user: state.user
    };
}

export default connect(mapStateToProps)(Sidebar);