import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:ui/common/svg_provider.dart';
import 'package:ui/models/agents.dart';
import 'package:ui/widgets/dot/index.dart';

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
      padding: EdgeInsets.symmetric(vertical: 15, horizontal: 10),
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
            height: 32,
            width: 32,
            margin: EdgeInsets.only(right: 20, left: 10,),
            decoration: BoxDecoration(
              image: DecorationImage(
                image: SvgProvider(
                'images/agents.svg',
                size: Size(128, 128),
                color: rgba(255, 255, 255, 1),
              ),
              fit: BoxFit.fill,
             ),
            ),
          ),
          getShowItem(children: [
            SelectableText(this.agent.name,
              textAlign: TextAlign.center,
              style: TextStyle(
              color: regularFontColor,
              fontSize: 12,
              fontWeight: FontWeight.w600,
            )),
            SizedBox(height: 2),
            Dot(
              dotTitle: this.agent.isOnline ? "Online".tr : "Offline".tr,
              dotColor: this.agent.isOnline ? onlineColor : offlineColor,
            ),
          ]),
          getShowItem(children: [
            SelectableText(this.agent.id,
              textAlign: TextAlign.center,
              style: TextStyle(
              color: regularFontColor,
              fontSize: 12,
              fontWeight: FontWeight.w600,
            )),
            SizedBox(height: 2),
            SelectableText("ID".tr,
              style: TextStyle(
              color: offlineColor,
              fontSize: 12,
              fontWeight: FontWeight.w400,
            ))
          ]),
          getShowItem(children: [
            SelectableText(this.agent.ip,
              textAlign: TextAlign.center,
              style: TextStyle(
              color: regularFontColor,
              fontSize: 12,
              fontWeight: FontWeight.w600,
            )),
            SizedBox(height: 2),
            SelectableText("IP".tr,
              style: TextStyle(
              color: offlineColor,
              fontSize: 12,
              fontWeight: FontWeight.w400,
            ))
          ]),
          getShowItem(children: [
            Row(
              children: [
                SelectableText(this.agent.networkSpeed?.toString() ?? "- -",
                  textAlign: TextAlign.end,
                  style: TextStyle(
                  color: regularFontColor,
                  fontSize: 24,
                  fontWeight: FontWeight.w600,
                  height: 1
                )),
                this.agent.networkSpeed != null ? SelectableText("  M/S",
                  textAlign: TextAlign.end,
                  style: TextStyle(
                  color: regularFontColor,
                  fontSize: 12,
                  fontWeight: FontWeight.w600,
                  height: 2
                )) : Container()
              ],
            ),
            SelectableText("Network speed".tr,
              style: TextStyle(
              color: offlineColor,
              fontSize: 12,
              fontWeight: FontWeight.w400,
            ))
          ]),
          getShowItem(children: [
            Row(
              children: [
                SelectableText(this.agent.taskNumber?.toString() ?? "- -",
                  textAlign: TextAlign.end,
                  style: TextStyle(
                  color: regularFontColor,
                  fontSize: 24,
                  fontWeight: FontWeight.w600,
                  height: 1
                )),
                this.agent.taskNumber != null ? SelectableText("  Tasks",
                  textAlign: TextAlign.end,
                  style: TextStyle(
                  color: regularFontColor,
                  fontSize: 12,
                  fontWeight: FontWeight.w600,
                  height: 2
                )) : Container()
              ],
            ),
            SelectableText("Running tasks".tr,
              style: TextStyle(
              color: offlineColor,
              fontSize: 12,
              fontWeight: FontWeight.w400,
            ))
          ])
        ],
      ),
    );
  }

  
}
