import React from 'react';
import {Route, IndexRoute} from 'react-router';

import App from '../App';
import GroupPage from '../Components/Groups/GroupPage';

export default (
    <Route path="/" component={App}>
        <Route path="group/:id" component={GroupPage}/>
    </Route>
)