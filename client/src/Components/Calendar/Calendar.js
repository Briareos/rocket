import React, {PropTypes} from "react";
import moment from "moment";

const Calendar = ({totalBodyCount, calendarData, warnings}) => {
    const cell = (month, cellDate) => {
        if (cellDate.format("MM") != month) {
            return (
                <div key={cellDate.format("YYYY-MM-DD")} className="calendar-cell"/>
            )
        }

        let dataForCell = calendarData[cellDate.format("YY-DD-MM")];

        return (
            <div key={cellDate.format("YYYY-MM-DD")} className="calendar-cell">
                <span className="calendar-header">{cellDate.format("D. MMM")}</span>
                <hr className="calendar-separator"/>
                <div className="calendar-cell-body">
                    <span>Available: {`${dataForCell && dataForCell.availableBodyCount}/${totalBodyCount}`}</span>
                    <br/>
                    {warnings && <span className="calendar-cell-warning">Warnings: {warnings}</span>}
                </div>
            </div>
        );

    };

    const cells = () => {
        const startOfMonth = moment().startOf('month');

        const activeMonth = startOfMonth.format("MM");

        let currentKey = startOfMonth;

        console.log(currentKey.format("d"));
        if (startOfMonth.format("d") != 0) {
            currentKey = startOfMonth.startOf('week');
        }

        var rows = [];
        var row = [];
        let rowId = 0;
        for (var i = 0; i < 42; i++) {
            if (i % 7 == 0 && i != 0) {
                rows.push(<div key={rowId} className="calendar-row">{row}</div>);
                rowId++;
                row = [];
            }

            row.push(cell(activeMonth, currentKey));
            currentKey.add('days', 1);
        }

        rows.push(<div key={rowId} className="calendar-row">{row}</div>);

        return rows;
    };


    let today = moment().format();

    console.log(calendarData);
    return (
        <div className="calendar">
            {cells()}
        </div>
    )
};

export default Calendar