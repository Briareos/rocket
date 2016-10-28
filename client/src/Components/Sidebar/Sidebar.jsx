import React, {Component} from "react";
import {connect} from "react-redux";
import * as groupSelectors from "../../selectors/groupSelectors";
import SidebarUserInfo from "./SidebarUserInfo";
import SidebarGroups from './SidebarGroups';

class Sidebar extends Component {
    constructor(props) {
        super(props);
    }

    render() {
        let {user, joinedGroups, watchedGroups} = this.props;
        console.log(user, joinedGroups, watchedGroups);
        return (
            <div className="sidebar">
                <SidebarUserInfo user={user}/>
                <SidebarGroups groups={joinedGroups} label={"Joined Groups"}/>
                <SidebarGroups groups={watchedGroups} label={"Watched Groups"}/>
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