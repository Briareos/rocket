import React from "react";
import {render} from "react-dom";
import {Router, browserHistory} from "react-router";
import {Provider} from "react-redux";
import routes from "./routes";
import configureStore from "./store/StoreConfiguration";
import {getProfile, getGroupDays} from "./actions/actionCreators";
import Api from './utils/Api';

const renderApplication = (initialState) => {
    const store = configureStore(initialState);

    render(
        <Provider store={store}>
            <Router history={browserHistory} routes={routes}/>
        </Provider>,
        document.getElementById('rocket-application')
    );
    document.getElementById('loading-state').className = "loading-state hide"
};

Api.get('profile').then(response => {
    let initialState = response.data;
    renderApplication(initialState);
});