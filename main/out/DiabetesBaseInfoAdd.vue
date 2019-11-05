<template>
    <v-container fluid v-if="formStruct">
        <material-card color="primary" title="基本信息">
            <v-row dense>
                <v-col cols="12" md="2" lg="2">
                    <v-text-field v-model="formData.name" label="姓名" dense></v-text-field>
                </v-col>
                <v-col cols="12" md="2" lg="2">
                    <v-text-field v-model="formData.age" label="年龄" dense></v-text-field>
                </v-col>
                <v-col cols="12" md="2" lg="2">
                    <v-select v-model="formData.sex" :items="formStruct.sex.data" label="性别" item-text="name"
                              item-value="_id" dense></v-select>
                </v-col>
                <v-col cols="12" md="2" lg="2">
                    <v-text-field v-model="formData.phoneNumber" label="手机号码" dense></v-text-field>
                </v-col>
                <v-col cols="12" md="2" lg="2">
                    <v-text-field v-model="formData.contactNumber" label="联系电话" dense></v-text-field>
                </v-col>
                <v-col cols="12" md="2" lg="2">
                    <v-text-field v-model="formData.relationshipWithPatients" label="与患者关系" dense></v-text-field>
                </v-col>
                <v-col cols="12" md="2" lg="2">
                    <v-text-field v-model="formData.registerDate" label="登记日期" dense></v-text-field>
                </v-col>
                <v-col cols="12" md="2" lg="2">
                    <v-text-field v-model="formData.registrant" label="登记人" dense></v-text-field>
                </v-col>
                <v-col cols="12" md="2" lg="2">
                    <v-select v-model="formData.diabetesType" :items="formStruct.diabetesType.data" label="糖尿病类型"
                              item-text="name" item-value="_id" dense></v-select>
                </v-col>
                <v-col cols="12" md="2" lg="2">
                    <v-text-field v-model="formData.placeOfAbode" label="通讯地址" dense></v-text-field>
                </v-col>
            </v-row>
        </material-card>
        <material-card color="primary" title="基本信息">
            <v-row dense>
                <v-col cols=6>
                    <v-radio-group
                            row
                            light
                            label="高血压"
                            dense
                            v-model="formData.highBloodPressure.value"
                            height="75"
                            class="radios"
                    >
                        <v-radio
                                dense color="primary"
                                v-for="n in formStruct.highBloodPressure.data"
                                :key="n._id"
                                :label="n.name"
                                :value="n"
                        ></v-radio>
                        <v-text-field v-model="formData.highBloodPressure.date" label="病龄" dense></v-text-field>
                    </v-radio-group>
                </v-col>
                <v-col cols=6>
                    <v-radio-group
                            row
                            light
                            label="糖尿病"
                            dense
                            v-model="formData.diabetes.value"
                            height="75"
                            class="radios"
                    >
                        <v-radio
                                dense color="primary"
                                v-for="n in formStruct.diabetes.data"
                                :key="n._id"
                                :label="n.name"
                                :value="n"
                        ></v-radio>
                        <v-text-field v-model="formData.diabetes.date" label="病龄" dense></v-text-field>
                    </v-radio-group>
                </v-col>
                <v-col cols=6>
                    <v-radio-group
                            row
                            light
                            label="心绞痛"
                            dense
                            v-model="formData.anginaPectoris.value"
                            height="75"
                            class="radios"
                    >
                        <v-radio
                                dense color="primary"
                                v-for="n in formStruct.anginaPectoris.data"
                                :key="n._id"
                                :label="n.name"
                                :value="n"
                        ></v-radio>
                        <v-text-field v-model="formData.anginaPectoris.date" label="病龄" dense></v-text-field>
                    </v-radio-group>
                </v-col>
                <v-col cols=6>
                    <v-radio-group
                            row
                            light
                            label="心肌梗塞"
                            dense
                            v-model="formData.myocardialInfarction.value"
                            height="75"
                            class="radios"
                    >
                        <v-radio
                                dense color="primary"
                                v-for="n in formStruct.myocardialInfarction.data"
                                :key="n._id"
                                :label="n.name"
                                :value="n"
                        ></v-radio>
                        <v-text-field v-model="formData.myocardialInfarction.date" label="病龄" dense></v-text-field>
                    </v-radio-group>
                </v-col>
                <v-col cols=6>
                    <v-radio-group
                            row
                            light
                            label="脑梗"
                            dense
                            v-model="formData.cerebralInfarction.value"
                            height="75"
                            class="radios"
                    >
                        <v-radio
                                dense color="primary"
                                v-for="n in formStruct.cerebralInfarction.data"
                                :key="n._id"
                                :label="n.name"
                                :value="n"
                        ></v-radio>
                        <v-text-field v-model="formData.cerebralInfarction.date" label="病龄" dense></v-text-field>
                    </v-radio-group>
                </v-col>
                <v-col cols=6>
                    <v-radio-group
                            row
                            light
                            label="其他"
                            dense
                            v-model="formData.other.value"
                            height="75"
                            class="radios"
                    >
                        <v-radio
                                dense color="primary"
                                v-for="n in formStruct.other.data"
                                :key="n._id"
                                :label="n.name"
                                :value="n"
                        ></v-radio>
                        <v-text-field v-model="formData.other.date" label="病龄" dense></v-text-field>
                    </v-radio-group>
                </v-col>
                <v-row dense>
                    <v-col cols="12" md="2" lg="2" class="symptom-box">非药物治疗</v-col>
                    <v-col cols="12" md="10" lg="10">
                        <v-row dense>
                            <v-col cols="12" md="12" lg="12" class="symptom">
                                <v-row dense>
                                    <v-checkbox
                                            dense
                                            class="mr-2"
                                            v-model="formData.nondrugTherapy"
                                            v-for="n in formStruct.nondrugTherapy.data" :key="n._id"
                                            :label="n.name"
                                            :value="n"
                                            color="primary"
                                            hide-details
                                    ></v-checkbox>
                                </v-row>
                            </v-col>
                        </v-row>
                    </v-col>
                </v-row>
                <v-col cols="12" md="2" lg="2">
                    <v-text-field v-model="formData.treatment" label="目前治疗情况" dense></v-text-field>
                </v-col>
                <v-col cols="12" md="2" lg="2">
                    <v-text-field v-model="formData.preparer" label="填表人" dense></v-text-field>
                </v-col>
                <v-col cols="12" md="2" lg="2">
                    <v-text-field v-model="formData.dateOfFilling" label="填表日期" dense></v-text-field>
                </v-col>
            </v-row>
        </material-card>

    </v-container>
</template>

<script>
    import {mapMutations} from 'vuex'

    export default {
        name: 'DiabetesBaseInfoAdd',
        data: () => (
            {
                "formData": {
                    "age": null,
                    "anginaPectoris": {"date": null, "value": null},
                    "cerebralInfarction": {"date": null, "value": null},
                    "contactNumber": null,
                    "dateOfFilling": null,
                    "diabetes": {"date": null, "value": null},
                    "diabetesType": null,
                    "highBloodPressure": {"date": null, "value": null},
                    "myocardialInfarction": {"date": null, "value": null},
                    "name": null,
                    "nondrugTherapy": [],
                    "other": {"date": null, "value": null},
                    "phoneNumber": null,
                    "placeOfAbode": null,
                    "preparer": null,
                    "registerDate": null,
                    "registrant": null,
                    "relationshipWithPatients": null,
                    "sex": null,
                    "treatment": null
                }, "formStruct": {}
            }
        ),
        methods: {
            ...mapMutations([
                'alert'
            ]),
            init() {
                const data = {key: 'DIABETES_BASE_INFO'}
                this.$http.get(this.$api.form.struct, data).then(({data}) => {
                    if (data && data.code === -1) {
                        this.$emit("alert", {
                            color: "red",
                            text: "表单结构获取失败！"
                        });
                        return;
                    }
                    // console.log(data);
                    this.formStruct = data.data.data;
                });
            },
        },
        created() {
            this.init()
        }
    }
</script>

<style scoped>
    .radios {
        border: 1px #ccc dashed;
        padding: 0 10px;
    }

    .radios /deep/ .v-input__control {
        height: 75px;
    }

    .radios /deep/ .v-text-field__slot {
        margin-top: 20px
    }
</style>
