import 'package:get/get.dart';
import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';

import '../../common/colors.dart';
import '../../models/service.dart';
import '../../widgets/more_actions/index.dart';
import '../../widgets/confirm/index.dart';
import '../../widgets/dotted_line/index.dart';

import 'controller.dart';

class ServicePanel extends GetView<ServiceController> {
  final Service service;

  ServicePanel({required this.service});

  @override
  Widget build(BuildContext context) {
    return Container(
      width: 312,
      height: 360,
      padding: EdgeInsets.only(top: 2, right: 2, bottom: 12, left: 12),
      decoration: BoxDecoration(
        color: Colors.white,
        borderRadius: BorderRadius.all(Radius.circular(6)),
        boxShadow: [
          BoxShadow(
            offset: Offset(0, 2),
            color: rgba(52, 61, 190, 0.1),
            blurRadius: 34,
          )
        ],
      ),
      child: Column(
        children: [
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            crossAxisAlignment: CrossAxisAlignment.center,
            children: [
              SelectableText(
                'Qingstor',
                style: TextStyle(
                  color: secondaryFontColor,
                  fontSize: 12,
                  height: 1.67,
                  fontWeight: FontWeight.w400,
                ),
              ),
              MoreActions(
                onSelected: (String option) {
                  Get.dialog(Confirm(
                      title: "Are you sure to delete this service ?".tr,
                      description:
                          "After deleting, it will not affect the created tasks, but it will not appear in the service list and the identity option of the created task."
                              .tr,
                      onConfirm: () {
                        controller.deleteIdentity(service).then((result) {
                          Get.back();
                        });
                      }));
                },
                itemBuilder: (BuildContext context) => [
                  PopupMenuItem(
                    value: "delete",
                    height: 32.0,
                    child: Text("Delete service".tr,
                        style: TextStyle(
                          fontSize: 12.0,
                          color: regularFontColor,
                        )),
                  ),
                ],
              ),
            ],
          ),
          SizedBox(height: 20),
          SvgPicture.asset(
            'images/qingstor_logo.svg',
            width: 32,
            height: 32,
          ),
          SizedBox(height: 12),
          SelectableText(
            "${service.type} - ${service.name}",
            style: TextStyle(
              fontSize: 12,
              fontWeight: FontWeight.w500,
              color: headlineFontColor,
              height: 1.67,
            ),
          ),
          SizedBox(height: 24),
          DottedLine(
            dotWidth: 3,
            color: lightLineColor,
          ),
          Container(
            padding: EdgeInsets.only(top: 24, right: 12, left: 12),
            alignment: Alignment.topLeft,
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                SelectableText(
                  'Credential',
                  style: TextStyle(
                    fontSize: 12,
                    height: 1.67,
                    color: disableFontColor,
                    fontWeight: FontWeight.w400,
                  ),
                ),
                SelectableText(
                  service.credential.protocol,
                  style: TextStyle(
                    fontSize: 12,
                    height: 1.67,
                    color: regularFontColor,
                    fontWeight: FontWeight.w500,
                  ),
                ),
                SelectableText(
                  service.credential.args[0],
                  style: TextStyle(
                    fontSize: 12,
                    height: 1.67,
                    color: regularFontColor,
                    fontWeight: FontWeight.w500,
                  ),
                ),
                SelectableText(
                  service.credential.args[1],
                  style: TextStyle(
                    fontSize: 12,
                    height: 1.67,
                    color: regularFontColor,
                    fontWeight: FontWeight.w500,
                  ),
                ),
                SizedBox(height: 26),
                SelectableText(
                  'Endpoint',
                  style: TextStyle(
                    fontSize: 12,
                    height: 1.67,
                    color: disableFontColor,
                    fontWeight: FontWeight.w400,
                  ),
                ),
                SelectableText(
                  "${service.endpoint.protocol}://${service.endpoint.host}:${service.endpoint.port}",
                  style: TextStyle(
                    fontSize: 12,
                    height: 1.67,
                    color: regularFontColor,
                    fontWeight: FontWeight.w500,
                  ),
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }
}
