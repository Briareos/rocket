import React from "react";
import {render} from "react-dom";
import {Router, browserHistory} from "react-router";
import {Provider} from "react-redux";
import routes from "./routes";
import configureStore from "./store/StoreConfiguration";
import Api from "./utils/Api";
import {getGroupDays} from "./actions/actionCreators";

const renderApplication = (initialState) => {
    const store = configureStore(initialState);
    store.dispatch(getGroupDays(1, 10, 2016)).then(() => console.log("Loaded"));
    render(
        <Provider store={store}>
            <Router history={browserHistory} routes={routes}/>
        </Provider>,
        document.getElementById('rocket-application')
    );
    document.getElementById('loading-state').className = "loading-state hide";
};

const renderLoginApp = (loginUrl) => {
    render(
        <LoginApp loginUrl={loginUrl}/>,
        document.getElementById('rocket-application')
    );
    document.getElementById('loading-state').className = "loading-state hide";
};

Api.get('profile').then(response => {
    if (response.data.error === "not_logged_in") {
        // window.location.pathname = browserHistory.push('login');
        renderLoginApp();
        return;
    }
    let initialState = response.data;
    renderApplication(initialState);
});