import moment from "moment";

export const getJoinedGroupsForUser = state => {
    return [...state.groups].filter(group => state.user.joined_groups.includes(group.id))
};

export const getWatchedGroupsForUser = state => {
    return [...state.groups].filter(group => state.user.watched_groups.includes(group.id))
};

export const getCalendarGroupData = state => {
    return state.groupCalendar.days.reduce((acc, day) => {
        acc[moment(day.Date).format("YY-DD-MM")] = day;

        return acc
    }, {});
};