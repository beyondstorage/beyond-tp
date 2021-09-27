import 'package:flutter/material.dart';
import 'package:ui/common/colors.dart';

import 'model.dart';
import 'tab_title.dart';

class Tabs extends StatefulWidget {

  final List<TabPane> titleList;

  const Tabs({
    required this.titleList,
  });
  
  @override
  _TabsState createState() => _TabsState();
}

class _TabsState extends State<Tabs> {
  late String selectString;
  late Widget showWidget;
  
  @override
  void initState() {
    super.initState();
    selectString = widget.titleList[0].tabTitle;
    showWidget = widget.titleList[0].pane;
  }

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Container(
            margin: EdgeInsets.all(10),
            decoration: BoxDecoration(
              color: Colors.white,
              boxShadow: [
                BoxShadow(offset: Offset(0, 1), color: rgba(226, 232, 240, 1)),
                BoxShadow(offset: Offset(-1, 0), color: Colors.white),
                BoxShadow(offset: Offset(1, 0), color: Colors.white),
              ],
            ),
            child: Row(
              children: widget.titleList
                .map((TabPane item) => TabTitle(
                  selected: selectString == item.tabTitle,
                  title: item.tabTitle,
                  onPressed: () {
                    setState(() {
                      selectString = item.tabTitle;
                      showWidget = item.pane;
                    });
                  })
                ).toList()
            ),
        ),
        Expanded(
          child: Container(
            child: showWidget,
          ),
        )
      ],
    );
  }
}
