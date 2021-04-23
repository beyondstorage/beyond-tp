import 'package:get/get.dart';
import 'package:flutter/material.dart';
import 'package:form_validator/form_validator.dart';

import './source_form_fields.dart';
import './destination_form_fields.dart';
import './controller.dart';

class CreateTaskForm extends GetView<CreateTaskController> {
  final GlobalKey<FormState> formKey;
  final Function onSubmit;

  CreateTaskForm(this.formKey, this.onSubmit);

  @override
  Widget build(BuildContext context) {
    return Form(
      key: formKey,
      autovalidateMode: controller.autoValidateMode.value,
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Container(
            padding: EdgeInsets.symmetric(horizontal: 30.0),
            margin: EdgeInsets.only(bottom: 32.0),
            child: Row(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Container(
                  width: 124,
                  height: 32,
                  alignment: Alignment.centerLeft,
                  child: RichText(
                    text: TextSpan(
                      text: 'Task name'.tr,
                      style: TextStyle(
                        fontSize: 12,
                        color: Theme.of(context).primaryColorLight,
                        fontWeight: FontWeight.w500,
                      ),
                      children: [
                        TextSpan(
                          text: '*',
                          style: TextStyle(
                            color: Colors.red,
                          ),
                        ),
                      ],
                    ),
                  ),
                ),
                Expanded(
                  child: TextFormField(
                    autofocus: true,
                    decoration: InputDecoration(
                      border: const OutlineInputBorder(),
                      isDense: true,
                      contentPadding:
                          EdgeInsets.symmetric(horizontal: 12, vertical: 14),
                    ),
                    maxLines: 1,
                    style: TextStyle(fontSize: 12),
                    textInputAction: TextInputAction.next,
                    keyboardType: TextInputType.text,
                    validator: ValidationBuilder()
                        .minLength(1, 'Please enter name')
                        .build(),
                    onSaved: controller.name,
                  ),
                ),
              ],
            ),
          ),
          Container(
            decoration: new BoxDecoration(
              border: Border(
                bottom: BorderSide(
                  style: BorderStyle.solid,
                  color: Color.fromRGBO(228, 235, 241, 1),
                ),
              ),
            ),
          ),
          SourceFormFields(),
          Container(
            decoration: new BoxDecoration(
              border: Border(
                bottom: BorderSide(
                  style: BorderStyle.solid,
                  color: Color.fromRGBO(228, 235, 241, 1),
                ),
              ),
            ),
          ),
          DestinationFormFields(),
        ],
      ),
    );
  }
}
