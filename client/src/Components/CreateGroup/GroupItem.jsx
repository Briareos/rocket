import React, {PropTypes} from 'react';

const GroupItem = ({group, ...props}) => (
    <div className="group-item">
        {group.name}
    </div>
);

GroupItem.propTypes = {
    group: PropTypes.object.isRequired
};

export default GroupItem;