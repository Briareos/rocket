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
        case actions.CREATE_GROUP:
            return Object.assign({}, state, {
                joined_groups: [...state.joined_groups, action.group.id],
                watched_groups: [...state.watched_groups, action.group.id],
            });
        default:
            return state;
    }
}