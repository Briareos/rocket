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

export function watchGroup(groupID) {
    return dispatch => {
        dispatch(watchGroupStarted());
        return Api.post('groupAction', {
            type: 'watch',
            groupID: groupID,
        }).then(
            response => {
                dispatch(watchGroupSuccess(groupID));
            },
            error => {
                dispatch(watchGroupFailure(error));
            }
        )
    }
}

function watchGroupStarted() {
    return {
        type: action.WATCH_GROUP,
    }
}

function watchGroupSuccess(groupID) {
    return {
        type: action.WATCH_GROUP_SUCCESS,
        groupID,
    }
}

function watchGroupFailure(message) {
    return {
        type: action.WATCH_GROUP_FAILURE,
        message,
    }
}

export function unwatchGroup(groupID) {
    return dispatch => {
        dispatch(unwatchGroupStarted());
        return Api.post('groupAction', {
            type: 'unwatch',
            groupID: groupID,
        }).then(
            response => {
                dispatch(unwatchGroupSuccess(groupID));
            },
            error => {
                dispatch(unwatchGroupFailure(error));
            }
        )
    }
}

function unwatchGroupStarted() {
    return {
        type: action.UNWATCH_GROUP,
    }
}

function unwatchGroupSuccess(groupID) {
    return {
        type: action.UNWATCH_GROUP_SUCCESS,
        groupID,
    }
}

function unwatchGroupFailure(message) {
    return {
        type: action.UNWATCH_GROUP_FAILURE,
        message,
    }
}

export function createRule(groupID, description, type, operation, threshold) {
    return dispatch => {
        dispatch(createRuleStarted());
        return Api.post('createRule', {
            groupID,
            rule: {
                description,
                type,
                operation,
                threshold,
            }
        }).then(
            response => {
                dispatch(createRuleSuccess(groupID, response.data.rule));
            },
            error => {
                dispatch(createRuleFailure(error));
            }
        )
    }
}

function createRuleStarted() {
    return {
        type: action.CREATE_RULE,
    }
}

function createRuleSuccess(groupID, rule) {
    return {
        type: action.CREATE_RULE_SUCCESS,
        groupID,
        rule,
    }
}

function createRuleFailure(message) {
    return {
        type: action.CREATE_RULE_FAILURE,
        message,
    }
}

export function muteRule(ruleID) {
    return dispatch => {
        dispatch(muteRuleStarted());
        return Api.post('ruleAction', {
            type: "mute",
            ruleID,
        }).then(
            () => {
                dispatch(muteRuleSuccess(ruleID));
            },
            error => {
                dispatch(muteRuleFailure(error));
            }
        )
    }
}

function muteRuleStarted() {
    return {
        type: action.MUTE_RULE,
    }
}

function muteRuleSuccess(ruleID) {
    return {
        type: action.MUTE_RULE_SUCCESS,
        ruleID,
    }
}

function muteRuleFailure(message) {
    return {
        type: action.MUTE_RULE_FAILURE,
        message,
    }
}

export function unmuteRule(ruleID) {
    return dispatch => {
        dispatch(unmuteRuleStarted());
        return Api.post('ruleAction', {
            type: "unmute",
            ruleID,
        }).then(
            () => {
                dispatch(unmuteRuleSuccess(ruleID));
            },
            error => {
                dispatch(unmuteRuleFailure(error));
            }
        )
    }
}

function unmuteRuleStarted() {
    return {
        type: action.UNMUTE_RULE,
    }
}

function unmuteRuleSuccess(ruleID) {
    return {
        type: action.UNMUTE_RULE_SUCCESS,
        ruleID,
    }
}

function unmuteRuleFailure(message) {
    return {
        type: action.UNMUTE_RULE_FAILURE,
        message,
    }
}