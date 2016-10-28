import React from 'react';
import {Route, IndexRoute} from 'react-router';

import App from '../App';
import GroupPage from '../Components/Groups/GroupPage';
import CreateGroupPage from '../Components/CreateGroup/CreateGroupPage';

export default (
    <Route path="/" component={App}>
        <Route path="group/:id" component={GroupPage}/>
        <Route path="group/create" component={CreateGroupPage}/>
    </Route>
)