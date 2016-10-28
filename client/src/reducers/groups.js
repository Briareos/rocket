import * as actions from "../actions/actions";

export default function (state = [], action) {
    switch (action.type) {
        case actions.GET_PROFILE_SUCCESS:
            return [...action.groups];
        case actions.CREATE_GROUP:
            return [...state, action.group];
        case actions.CREATE_RULE:
            return [...state].map(group => {
                if (group.id == action.groupID) {
                    group.rules = [...group.rules, action.rule];
                }

                return group;
            });
        default:
            return state;
    }
}