import 'package:flutter/material.dart';
import 'package:ui/common/colors.dart';
import 'package:ui/models/task.dart';

class TimeLogPane extends StatelessWidget {
  
  final String logTitle;
  final String description;
  final List<TimeLog> logList;

  const TimeLogPane({
    required this.logTitle,
    required this.description,
    required this.logList
  });
  

  @override
  Widget build(BuildContext context) {
    return Container( 
      color: timeLogBackgroundColor,
      padding: EdgeInsets.symmetric(vertical: 20),
      child: Column(
        children: [
          Row(
            children: [
              Expanded(
                child: Container(
                  padding: EdgeInsets.only(left: 20, bottom: 20),
                  decoration: BoxDecoration(
                    color: timeLogBackgroundColor,
                    boxShadow: [
                      BoxShadow(offset: Offset(0, 1), color: rgba(226, 232, 240, 0.2))
                    ],
                  ),
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      SelectableText(
                        logTitle,
                        style: TextStyle(
                          color: rgba(202, 209, 216, 1),
                          fontSize: 18,
                          fontWeight: FontWeight.w600,
                          fontFamily: 'Roboto'
                        ),
                      ),
                      SelectableText(
                        description,
                        style: TextStyle(
                          color: logFontColor,
                          fontSize: 12,
                          fontWeight: FontWeight.normal,
                          fontFamily: 'Roboto'
                        ),
                      )
                    ],
                  )
                ),
              )
            ],
          ),
          Expanded(
            child: ListView(
              children: logList.length > 0
                ? logList.map((TimeLog item) => Container(
                    padding: EdgeInsets.only(left: 20),
                    child: Row(
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: [
                        Row(
                          children: [
                            Container(
                              padding: EdgeInsets.only(top: 10, bottom: 10, right: 7),
                              margin: EdgeInsets.only(top: 3),
                              child: Icon(
                                Icons.check_circle,
                                size: 16,
                                color: logFontColor,
                              ),
                            ),
                            SelectableText(
                              item.logContent,
                              style: TextStyle(
                                color: logFontColor,
                                fontSize: 12,
                                fontWeight: FontWeight.normal
                              )
                            )
                          ],
                        ),
                        Padding(
                          padding: EdgeInsets.only(right: 10),
                          child: SelectableText(
                            '${item.time}s',
                            style: TextStyle(
                              color: logFontColor,
                              fontSize: 12,
                              fontWeight: FontWeight.normal,
                              fontFamily: 'Roboto'
                            )
                          ),
                        )
                      ],
                    ),
                )).toList()
                : [],
            ),
          )
        ],
      ),
    );
  }
}
