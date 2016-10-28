import Api from "../utils/Api";
import * as action from "./actions";

export function getProfile() {
    return dispatch => {
        dispatch(getProfileStarted());
        return Api.get('profile').then(
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
        return Api.get('groupDays', {
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
        return Api.post('groupCreate', {
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
        return Api.post('groupAction', {
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


export function leaveGroup(groupID) {
    return dispatch => {
        dispatch(leaveGroupStarted());
        return Api.post('groupAction', {
            type: 'leave',
            groupID: groupID,
        }).then(
            response => {
                dispatch(leaveGroupSuccess(groupID));
            },
            error => {
                dispatch(leaveGroupFailure(error));
            }
        )
    }
}

function leaveGroupStarted() {
    return {
        type: action.LEAVE_GROUP,
    }
}

function leaveGroupSuccess(groupID) {
    return {
        type: action.LEAVE_GROUP_SUCCESS,
        groupID,
    }
}

function leaveGroupFailure(message) {
    return {
        type: action.LEAVE_GROUP_FAILURE,
        message,
    }
}


//ruleCreate, groupAction, ruleAction (primaju type)