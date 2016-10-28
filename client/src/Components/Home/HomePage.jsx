import React, {Component} from 'react';
import {connect} from 'react-redux';

class HomePage extends Component {
    constructor(props) {
        super(props);
    }
    render() {
        return (
            <div id="homeage" className="page" >
                home page
            </div>
        );
    }
}

function mapStateToProps (state) {
    return {
        user: state.user
    };
}

export default connect(mapStateToProps)(HomePage);