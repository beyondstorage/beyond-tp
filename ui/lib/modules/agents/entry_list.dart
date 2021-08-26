import 'package:get/get.dart';
import 'package:flutter/material.dart';
import 'package:ui/modules/agents/agent_item.dart';
import 'package:ui/widgets/message/index.dart';
import 'package:ui/widgets/page_container/index.dart';

import 'controller.dart';
import 'toolbar.dart';

class EntryList extends GetView<AgentsController> {
  @override
  Widget build(BuildContext context) {
    return Expanded(
      child: Obx(() => WidgetContainer(
        child: Container(
          child: Padding(
            padding: EdgeInsets.all(20),
            child: Column(
              mainAxisAlignment: MainAxisAlignment.start,
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Toolbar(
                  onClick: () {
                    controller.getAgents(callBack: () {
                      Message.success(
                        context: Get.overlayContext as BuildContext,
                        message: "Refresh successfully！",
                      );
                    });
                  },
                ),
                Expanded(
                  flex: 1,
                  child: ListView(
                    shrinkWrap: true,
                    children: [
                      ...controller.agents.value.agents.map((agent) => AgentItem(agent: agent))
                    ],
                  ),
                ),
              ],
            ),
          ),
        ),
      )),
    );
  }
}
