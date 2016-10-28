import {combineReducers} from "redux";
import user from "./user";
import groups from "./groups";
import users from "./users";
import groupCalendar from "./groupCalendar";

const reducer = combineReducers({
    user,
    groups,
    users,
    groupCalendar,
});

export default reducer;