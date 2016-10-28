import React, {PropTypes} from 'react';
import classnames from 'classnames';

const GroupItem = ({group, selected, ...props}) => (
    <div className={classnames("group-item", {'selected': selected})} {...props}>
        {group.name}
    </div>
);

GroupItem.propTypes = {
    group: PropTypes.object.isRequired,
    selected: PropTypes.bool.isRequired
};

export default GroupItem;