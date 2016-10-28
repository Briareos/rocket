import React, {Component, PropTypes} from "react";

import Sidebar from './Components/Sidebar';
import Header from './Components/Header';

class App extends Component {
    constructor (props) {
        super(props);
    }
    render () {
        return (
            <div className="application-wrapper">
                <Sidebar/>
                <div className="main-container">
                    <Header/>
                    <div className="content-wrapper">
                        {this.props.children}
                    </div>
                </div>

            </div>
        )
    }
}

App.propTypes = {
    children: PropTypes.object
};

export default App;