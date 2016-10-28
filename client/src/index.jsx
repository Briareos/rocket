import React from 'react';
import {render} from 'react-dom';
import {Router, browserHistory} from 'react-router';
import {Provider} from 'react-redux';

import routes from './routes';
import configureStore from './store/StoreConfiguration';

const store = configureStore();

render(
    <Provider store={store}>
        <Router history={browserHistory} routes={routes}/>
    </Provider>,
    document.getElementById('luthor-application')
);