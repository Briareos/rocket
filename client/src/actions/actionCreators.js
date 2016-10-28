import Api from "../utils/Api";
import * as action from "./actions";

export function getProfile() {
    return dispatch => {
        dispatch(getProfileStarted());
        console.log(getProfileStarted());
        Api.get('profile').then(
            response => {
                console.log(getProfileSuccess(response.data.user, response.data.groups, response.data.users));
                dispatch(getProfileSuccess(response.data.user, response.data.groups, response.data.users));
            },
            error => {
                console.log(getProfileFailure(error));
                dispatch(getProfileFailure(error));
            }
        )
    }
}

function getProfileStarted() {
    return {
        type: action.GET_PROFILE
    }
}

function getProfileSuccess(user, groups, users) {
    return {
        type: action.GET_PROFILE_SUCCESS,
        user,
        groups,
        users,
    }
}

function getProfileFailure(message) {
    return {
        type: action.GET_PROFILE_FAILURE,
        message,
    }
}

