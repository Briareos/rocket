import React, {Component} from 'react';
import {connect} from 'react-redux';

class GroupPage extends Component {
    constructor(props) {
        super(props);
    }
    render() {
        return (
            <div id="group-page" className="page" >
                group page
            </div>
        );
    }
}

function mapStateToProps (state) {
    return {
        groups: state.groups
    };
}

export default connect(mapStateToProps)(GroupPage);