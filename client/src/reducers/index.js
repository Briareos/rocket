import {combineReducers} from "redux";
import user from "./user";
import groups from "./groups";
import users from "./users";

const reducer = combineReducers({
    user,
    groups,
    users,
});

export default reducer;