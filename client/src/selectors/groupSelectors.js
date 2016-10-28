export const getJoinedGroupsForUser = state => {
    return [...state.groups].filter(group => state.user.joined_groups.includes(group.id))
};

export const getWatchedGroupsForUser = state => {
    return [...state.groups].filter(group => state.user.watched_groups.includes(group.id))
};