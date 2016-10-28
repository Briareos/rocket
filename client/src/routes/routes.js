import React from 'react';
import {Route, IndexRoute} from 'react-router';

import App from '../App';
import GroupPage from '../Components/Groups/GroupPage';
import CreateGroupPage from '../Components/CreateGroup/CreateGroupPage';
import HomePage from '../Components/Home/HomePage';
import LoginPage from '../Components/Login/LoginPage';

export default (
    <Route path="/" component={App}>
        <Route path="home" component={HomePage}/>
        <Route path="login" component={LoginPage}/>
        <Route path="group/create" component={CreateGroupPage}/>
        <Route path="group/:id" component={GroupPage}/>
    </Route>
)