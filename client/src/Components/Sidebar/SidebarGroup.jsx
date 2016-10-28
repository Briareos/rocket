import React, {PropTypes} from 'react';

const SidebarGroup = ({group, ...props}) => (
    <div className="sidebar-group">
        {group.name}
    </div>
);

SidebarGroup.propTypes = {
    group: PropTypes.object.isRequired
};

export default SidebarGroup;