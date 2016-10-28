import React, {Component} from 'react';

class LoginApp extends Component {
    constructor(props) {
        super(props);
    }
    render() {
        return (
            <div className="login-form">
                <button>Login with Google</button>
            </div>
        )
    }
}

export default LoginApp;