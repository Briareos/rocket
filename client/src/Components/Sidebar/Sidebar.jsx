import React, {Component} from "react";
import {connect} from "react-redux";
import * as groupSelectors from "../../selectors/groupSelectors";
import SidebarUserInfo from "./SidebarUserInfo";

class Sidebar extends Component {
    constructor(props) {
        super(props);
    }

    render() {
        console.log(this.props);
        return (
            <div className="sidebar">
                <SidebarUserInfo user={this.props.user}/>
            </div>
        );
    }
}

function mapStateToProps(state) {
    return {
        user: state.user,
        joinedGroups: groupSelectors.getJoinedGroupsForUser(state),
        watchedGroups: groupSelectors.getWatchedGroupsForUser(state)
    };
}

export default connect(mapStateToProps)(Sidebar);