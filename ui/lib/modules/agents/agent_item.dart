import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:ui/models/agents.dart';

import '../../common/colors.dart';
import 'controller.dart';

class AgentItem extends GetView<AgentsController> {
  final Agent agent;

  AgentItem({required this.agent});

  Widget getShowItem({int flex = 1, required List<Widget> children}) {
    return Expanded(
        flex: 1,
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          mainAxisAlignment: MainAxisAlignment.center,
          children: children,
        ));
  }
  
  @override
  Widget build(BuildContext context) {
    return Container(
      padding: EdgeInsets.symmetric(vertical: 10, horizontal: 10),
      margin: EdgeInsets.only(bottom: 20),
      decoration: BoxDecoration(
        border: Border.all(width: 1, color: lightLineColor),
        borderRadius: BorderRadius.all(Radius.circular(6.0)),
      ),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.start,
        crossAxisAlignment: CrossAxisAlignment.center,
        children: [
          Container(
            height: 50,
            width: 50,
            margin: EdgeInsets.only(right: 16.0),
            child: Icon(
              Icons.dns_sharp,
              color: regularLineColor,
              size: 32,
            ),
          ),
          getShowItem(children: [
            SelectableText(this.agent.name,
              style: TextStyle(
              color: regularFontColor,
              fontSize: 12,
              fontWeight: FontWeight.w600,
            )),
            SizedBox(height: 2),
            Row(
              mainAxisAlignment: MainAxisAlignment.start,
              crossAxisAlignment: CrossAxisAlignment.center,
              children: [
                Container(
                  height: 8,
                  width: 8,
                  margin: EdgeInsets.only(right: 4, top: 3),
                  decoration: BoxDecoration(
                    borderRadius: BorderRadius.all(Radius.circular(4)),
                    color: this.agent.isOnline ? onlineColor : offlineColor
                  ),
                ),
                SelectableText(this.agent.isOnline ? "Online".tr : "Offline".tr,
                  style: TextStyle(
                  color: offlineColor,
                  fontSize: 12,
                  fontWeight: FontWeight.w400,
                ))
              ],
            )
          ]),
          getShowItem(children: [
            SelectableText(this.agent.id,
              style: TextStyle(
              color: regularFontColor,
              fontSize: 12,
              fontWeight: FontWeight.w600,
            )),
            SizedBox(height: 2),
            SelectableText("ID".tr,
              style: TextStyle(
              color: regularFontColor,
              fontSize: 12,
              fontWeight: FontWeight.w400,
            ))
          ]),
          getShowItem(children: [
            SelectableText(this.agent.ip,
              style: TextStyle(
              color: regularFontColor,
              fontSize: 12,
              fontWeight: FontWeight.w600,
            )),
            SizedBox(height: 2),
            SelectableText("IP".tr,
              style: TextStyle(
              color: regularFontColor,
              fontSize: 12,
              fontWeight: FontWeight.w400,
            ))
          ]),
          getShowItem(children: [
            Row(
              crossAxisAlignment: CrossAxisAlignment.end,
              children: [
                SelectableText(this.agent.networkSpeed?.toString() ?? "- -",
                  style: TextStyle(
                  color: regularFontColor,
                  fontSize: 24,
                  fontWeight: FontWeight.w600,
                )),
                this.agent.networkSpeed != null ? SelectableText("  M/S",
                  style: TextStyle(
                  color: regularFontColor,
                  fontSize: 12,
                  fontWeight: FontWeight.w600,
                )) : Container()
              ],
            ),
            SelectableText("Network speed".tr,
              style: TextStyle(
              color: regularFontColor,
              fontSize: 12,
              fontWeight: FontWeight.w400,
            ))
          ]),
          getShowItem(children: [
            Row(
              mainAxisAlignment: MainAxisAlignment.start,
              crossAxisAlignment: CrossAxisAlignment.end,
              children: [
                SelectableText(this.agent.taskNumber?.toString() ?? "- -",
                  style: TextStyle(
                  color: regularFontColor,
                  fontSize: 24,
                  fontWeight: FontWeight.w600,
                )),
                this.agent.taskNumber != null ? SelectableText("  Tasks",
                  style: TextStyle(
                  color: regularFontColor,
                  fontSize: 12,
                  fontWeight: FontWeight.w600,
                )) : Container()
              ],
            ),
            SelectableText("Running tasks".tr,
              style: TextStyle(
              color: regularFontColor,
              fontSize: 12,
              fontWeight: FontWeight.w400,
            ))
          ])
        ],
      ),
    );
  }

  
}
