import React from 'react';
import {render} from 'react-dom';
import {Router, browserHistory} from 'react-router';
import {Provider} from 'react-redux';
import Axios from 'axios';

import routes from './routes';
import configureStore from './store/StoreConfiguration';

const renderApplication = (initialStore) => {
    const store = configureStore(initialStore);

    render(
        <Provider store={store}>
            <Router history={browserHistory} routes={routes}/>
        </Provider>,
        document.getElementById('rocket-application')
    );
};

