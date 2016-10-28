import React, {PropTypes} from 'react';

import SidebarGroup from './SidebarGroup';

const SidebarGroups = ({groups, label, ...props}) => (
    <div className="sidebar-groups">
        <div className="groups-label">
            {label}
        </div>
        <div className="groups-wrapper">
            {groups.map(group => <SidebarGroup group={group} key={group.id}/>)}
        </div>
    </div>
);

SidebarGroups.propTypes = {
    label: PropTypes.string.isRequired,
    groups: PropTypes.array.isRequired
};

export default SidebarGroups;