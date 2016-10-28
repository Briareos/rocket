import React from 'react';
import {render} from 'react-dom';
import {Router, browserHistory} from 'react-router';
import {Provider} from 'react-redux';

import routes from './routes';
import configureStore from './store/StoreConfiguration';
import Api from './utils/Api';

const renderApplication = (initialStore) => {
    const store = configureStore(initialStore);

    render(
        <Provider store={store}>
            <Router history={browserHistory} routes={routes}/>
        </Provider>,
        document.getElementById('rocket-application')
    );
};

Api.get('profile').then(response => {
    console.log(response.data);
});