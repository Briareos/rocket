import * as actions from "../actions/actions";

export default function (state = [], action) {
    switch (action.type) {
        case actions.GET_PROFILE_SUCCESS:
            return [...action.users];
        default:
            return state;
    }
}