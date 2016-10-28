import React, {Component} from "react";
import {connect} from "react-redux";
import * as selectors from "../../selectors/groupSelectors";
import Calendar from "../Calendar/Calendar";

class GroupPage extends Component {
    constructor(props) {
        super(props);
    }

    render() {
        return (
            <div id="group-page" className="page" >
                <Calendar calendarData={this.props.calendarData} totalBodyCount={this.props.totalBodyCount}/>
            </div>
        );
    }
}

function mapStateToProps (state) {
    return {
        group: state.group,
        calendarData: selectors.getCalendarGroupData(state),
        totalBodyCount: state.groupCalendar.bodyCountRightNow,
    };
}

export default connect(mapStateToProps)(GroupPage);