import Api from "../utils/Api";
import * as action from "./actions";

export function getProfile() {
    return dispatch => {
        dispatch(getProfileStarted());
        Api.get('profile').then(
            response => {
                dispatch(getProfileSuccess(response.data.user, response.data.groups, response.data.users));
            },
            error => {
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

export function getGroupDays(id, month, year) {
    return dispatch => {
        dispatch(getGroupDaysStarted());
        Api.get('groupDays', {
            params: {id, month, year}
        }).then(
            response => {
                dispatch(getGroupDaysSuccess(id, response.data.Days, response.data.totalBodyCount));
            },
            error => {
                dispatch(getGroupDaysFailure(error));
            }
        )
    }
}

function getGroupDaysStarted() {
    return {
        type: action.GET_GROUP_DAYS
    }
}

function getGroupDaysSuccess(groupID, days, bodyCount) {
    return {
        type: action.GET_GROUP_DAYS_SUCCESS,
        groupID,
        days,
        bodyCount,
    }
}

function getGroupDaysFailure(message) {
    return {
        type: action.GET_GROUP_DAYS_FAILURE,
        message,
    }
}

export function createGroup(name, description, busyValue = false, remoteValue = true) {
    return dispatch => {
        dispatch(createGroupStarted());
        Api.post('groupCreate', {
            name,
            description,
            busyValue,
            remoteValue,
        }).then(
            response => {
                dispatch(createGroupSuccess(response.data));
            },
            error => {
                dispatch(createGroupFailure(error));
            }
        )
    }
}

function createGroupStarted() {
    return {
        type: action.CREATE_GROUP
    }
}

function createGroupSuccess(group) {
    return {
        type: action.CREATE_GROUP_SUCCESS,
        group,
    }
}

function createGroupFailure(message) {
    return {
        type: action.CREATE_GROUP_FAILURE,
        message,
    }
}

export function joinGroup(groupID) {
    return dispatch => {
        dispatch(joinGroupStarted());
        Api.post('groupAction', {
            type: 'join',
            groupID: groupID,
        }).then(
            response => {
                dispatch(joinGroupSuccess(groupID));
            },
            error => {
                dispatch(joinGroupFailure(error));
            }
        )
    }
}

function joinGroupStarted() {
    return {
        type: action.JOIN_GROUP,
    }
}

function joinGroupSuccess(groupID) {
    return {
        type: action.JOIN_GROUP_SUCCESS,
        groupID,
    }
}

function joinGroupFailure(message) {
    return {
        type: action.JOIN_GROUP_FAILURE,
        message,
    }
}


//ruleCreate, groupAction, ruleAction (primaju type)