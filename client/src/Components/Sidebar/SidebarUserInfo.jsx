import React, {PropTypes} from 'react';

const SidebarUserInfo = ({user, ...props}) => (
    <div className="sidebar-user-info">
        {user.first_name} {user.last_name}
    </div>
);

SidebarUserInfo.propTypes = {
    user: PropTypes.object.isRequired
};

export default SidebarUserInfo;