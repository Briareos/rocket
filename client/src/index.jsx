import React from "react";
import {render} from "react-dom";
import {Router, browserHistory} from "react-router";
import {Provider} from "react-redux";
import routes from "./routes";
import configureStore from "./store/StoreConfiguration";
import {getProfile, getGroupDays} from "./actions/actionCreators";

const renderApplication = (initialStore) => {
    const store = configureStore(initialStore);
    store.dispatch(getProfile());
    store.dispatch(getGroupDays(1, 10, 2016));

    render(
        <Provider store={store}>
            <Router history={browserHistory} routes={routes}/>
        </Provider>,
        document.getElementById('rocket-application')
    );
};

renderApplication();