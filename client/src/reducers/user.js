import * as actions from "../actions/actions";

export default function (state = {
    statuses: [],
    joined_groups: [],
    watched_groups: [],
    muted_rules: [],
}, action) {
    switch (action.type) {
        case actions.GET_PROFILE_SUCCESS:
            return Object.assign({}, state, action.user);
        case actions.CREATE_GROUP_SUCCESS:
            return Object.assign({}, state, {
                joined_groups: [...state.joined_groups, action.group.id],
                watched_groups: [...state.watched_groups, action.group.id],
            });
        case actions.JOIN_GROUP_SUCCESS:
            return Object.assign({}, state, {
                joined_groups: [...state.joined_groups, action.groupID],
                watched_groups: [...state.watched_groups, action.groupID],
            });
        case actions.LEAVE_GROUP_SUCCESS:
            return Object.assign({}, state, {
                joined_groups: [...state.joined_groups].filter(groupID => groupID != action.groupID),
                watched_groups: [...state.watched_groups].filter(groupID => groupID != action.groupID),
            });
        case actions.WATCH_GROUP_SUCCESS:
            return Object.assign({}, state, {
                watched_groups: [...state.watched_groups, action.groupID],
            });
        case actions.UNWATCH_GROUP_SUCCESS:
            return Object.assign({}, state, {
                watched_groups: [...state.watched_groups].filter(groupID => groupID != action.groupID),
            });
        case actions.MUTE_RULE_SUCCESS:
            return Object.assign({}, state, {
                muted_rules: [...state.muted_rules, action.ruleID],
            });
        case actions.UNMUTE_RULE_SUCCESS:
            return Object.assign({}, state, {
                muted_rules: [...state.muted_rules].filter(ruleID => ruleID != action.ruleID),
            });
        default:
            return state;
    }
}