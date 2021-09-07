import 'package:get/get.dart';
import 'package:flutter/material.dart';
import 'package:ui/widgets/page_container/index.dart';
import '../../common/colors.dart';
import 'controller.dart';

class Toolbar extends GetView<AgentsController> {
  final Function onClick;

  Toolbar({required this.onClick});

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: EdgeInsets.only(bottom: 15),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.start,
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Container(
            child: GestureDetector(
              onTap: () => this.onClick(),
              child: Icon(
                IconData(0xe604, fontFamily: 'tpIcon'),
                size: 20,
                color: Colors.black87,
              ),
            ),
          ),
          SizedBox(
            width: 17,
          ),
          Obx(() => SelectableText(
                '${controller.agents.value.length()} ${controller.agents.value.length() > 1 ? 'Agents' : 'Agent'}',
                style: TextStyle(
                  color: regularFontColor,
                  fontSize: 12,
                  fontWeight: FontWeight.w400,
                ),
              ))
        ],
      ),
    );
  }
}
