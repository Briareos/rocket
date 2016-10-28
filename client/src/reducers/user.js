import * as actions from "../actions/actions";

export default function (state = {}, action) {
    switch (action.type) {
        case actions.GET_PROFILE_SUCCESS:
            return Object.assign({}, state, action.user);
        default:
            return state;
    }
}