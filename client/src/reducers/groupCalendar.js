import * as actions from "../actions/actions";

export default function (state = {
    groupID: null,
    days: [],
    bodyCountRightNow: 0,
}, action) {
    switch (action.type) {
        case actions.GET_GROUP_DAYS_SUCCESS:
            return Object.assign({}, state, {
                groupID: action.groupID,
                days: action.days,
                bodyCountRightNow: action.bodyCount,
            });
        default:
            return state;
    }
}