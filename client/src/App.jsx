import React, {Component, PropTypes} from 'react';

class App extends Component {
    constructor (props) {
        super(props)
    }
    render () {
        return (
            <div>
                <nav>header</nav>
                {this.props.children}
            </div>
        )
    }
}

App.propTypes = {
    children: PropTypes.object.isRequired
};

export default App;