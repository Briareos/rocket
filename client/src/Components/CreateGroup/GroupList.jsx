import React from 'react';

const GroupList = ({children, ...props}) => (
    <div className="group-list" {...props}>
        {children}
    </div>
);

export default GroupList;