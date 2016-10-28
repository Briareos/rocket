import React, {Component} from 'react';
import {connect} from 'react-redux';

class LoginPage extends Component {
    constructor(props) {
        super(props);
    }
    render() {
        return (
            <div id="login-page" className="page" >
                login page
            </div>
        );
    }
}

function mapStateToProps (state) {
    return {
        user: state.user
    };
}

export default connect(mapStateToProps)(LoginPage);