<template>
  <v-container fluid v-if="formStruct">
    <material-card color="primary" title="基础信息">

      <v-row dense>
        <v-col cols="12" md="6" lg="6" class="pa-0">
          <region-of-china ref="region" @getRegion="getRegion"></region-of-china>
        </v-col>
        <v-col cols="12" md="2" lg="2">
          <v-text-field v-model="formData.householder" label="户主姓名" dense></v-text-field>
        </v-col>
        <v-col cols="12" md="2" lg="2">
          <v-text-field v-model="formData.contactNumber" label="联系电话" dense></v-text-field>
        </v-col>
        <v-col cols="12" md="2" lg="2">
          <v-text-field v-model="formData.phoneNumber" label="手机号码" dense></v-text-field>
        </v-col>
        <v-col cols="12" md="2" lg="2">
          <v-text-field v-model="formData.familyAddress" label="家庭住址" dense></v-text-field>
        </v-col>
        <v-col cols="12" md="2" lg="2">
          <v-select v-model="formData.familyType" :items="formStruct.familyType" label="家庭类型" dense></v-select>
        </v-col>
        <v-col cols="12" md="2" lg="2">
          <v-text-field v-model="formData.residenceBookletCode" label="户口簿编号" dense></v-text-field>
        </v-col>
        <v-col cols="12" md="2" lg="2">
          <v-text-field v-model="formData.community" label="居住小区" dense></v-text-field>
        </v-col>
      </v-row>
    </material-card>

    <material-card color="primary" title="居住条件与卫生设施">
      <v-row>
        <v-col cols="12" md="2" lg="2">
          <v-select v-model="formData.housingType" :items="formStruct.housingType" label="住房类型" dense></v-select>
        </v-col>
        <v-col cols="12" md="2" lg="2">
          <v-text-field v-model="formData.housingArea" label="住房面积" dense></v-text-field>
        </v-col>
        <v-col cols="12" md="2" lg="2">
          <v-select v-model="formData.wcType" :items="formStruct.wcType" label="厕所" dense></v-select>
        </v-col>
        <v-col cols="12" md="2" lg="2">
          <v-select v-model="formData.fuelType" :items="formStruct.fuelType" label="燃料类型" dense></v-select>
        </v-col>
        <v-col cols="12" md="2" lg="2">
          <v-select v-model="formData.monthlyPerCapitaIncome" :items="formStruct.monthlyPerCapitaIncome" label="月人均收入"
                    dense></v-select>
        </v-col>
        <v-col cols="12" md="2" lg="2">
          <v-select v-model="formData.drinkingWater" :items="formStruct.drinkingWater" label="饮水" dense></v-select>
        </v-col>
        <v-col cols="12" md="2" lg="2">
          <v-select v-model="formData.kitchenExhaustType" :items="formStruct.kitchenExhaustType" label="厨房排风设施"
                    dense></v-select>
        </v-col>
        <v-col cols="12" md="2" lg="2">
          <v-select v-model="formData.airState" :items="formStruct.airState" label="温度情况" dense></v-select>
        </v-col>
        <v-col cols="12" md="2" lg="2">
          <v-select v-model="formData.lightingState" :items="formStruct.lightingState" label="采光情况" dense></v-select>
        </v-col>
        <v-col cols="12" md="2" lg="2">
          <v-select v-model="formData.warmAndCold" :items="formStruct.warmAndCold" label="保暖情况" dense></v-select>
        </v-col>
        <v-col cols="12" md="2" lg="2">
          <v-select v-model="formData.healthState" :items="formStruct.healthState" label="卫生情况" dense></v-select>
        </v-col>
        <v-col cols="12" md="2" lg="2">
          <v-select v-model="formData.poultryPen" :items="formStruct.poultryPen" label="禽畜栏" dense></v-select>
        </v-col>
      </v-row>
    </material-card>


    <material-card color="primary" title="家庭成员列表">
      <v-row dense>
        <v-spacer/>
        <v-col lg="6" align="end">
          <v-btn class="ml-2" outlined color="primary">
            <v-icon left>mdi-delete</v-icon>
            删除
          </v-btn>
          <v-btn class="ml-2" color="primary">
            <v-icon left>mdi-plus</v-icon>
            新增
          </v-btn>
        </v-col>
      </v-row>
      <v-data-table :headers="headers" :items="item" hide-default-footer locale="zh-cn"
                    no-data-text="暂无数据"/>
      <v-row dense>
        <v-col cols="12" md="6" lg="6"></v-col>
        <v-col cols="12" md="2" lg="2">
          <DoctorSelect @change="founderCode" label="建档人"/>
        </v-col>
        <v-col cols="12" md="2" lg="2">
          <DoctorSelect @change="personLiableCode" label="责任人"/>
        </v-col>
        <v-col cols="12" md="2" lg="2">
          <v-menu
            v-model="menu3"
            :close-on-content-click="false"
            :nudge-right="40"
            transition="scale-transition"
            offset-y
            min-width="290px"
          >
            <template dense v-slot:activator="{ on }">
              <v-text-field dense label="日期" v-model="date" v-on="on"></v-text-field>
            </template>
            <v-date-picker
              dense
              :first-day-of-week="0"
              locale="zh-cn"
              year-icon="mdi-calendar-blank"
              @input="menu3 = false"
              v-model="date"
              @change="profileDate"
            ></v-date-picker>
          </v-menu>
        </v-col>
        <v-col cols="12" md="3" lg="3"></v-col>
        <v-col cols="12" md="3" lg="3" align="center">
          <v-btn class="ml-2" color="primary" width="150" @click="save">
            保存
          </v-btn>

        </v-col>
        <v-col cols="12" md="3" lg="3" align="center">
          <v-btn class="ml-2" outlined color="primary" width="150" @click.stop="reset">
            重置
          </v-btn>
        </v-col>
        <v-col cols="12" md="3" lg="3"></v-col>
      </v-row>
    </material-card>
  </v-container>
</template>

<script>
    import { mapMutations } from 'vuex'

    export default {
        name: 'FamilyProfileInfoAdd',
        data: () => ({
            sex: [{ text: '男', value: 1 }, { text: '女', value: 2 }],
            date1: '',
            menu1: false,
            date2: '',
            menu2: false,
            date3: '',
            menu3: false,
            radioGroup: 1,
            formStruct: null,
            date: null,
            headers: [
                {
                    sortable: false,
                    text: '序号',
                    value: 'id'
                },
                {
                    sortable: false,
                    text: '成员姓名',
                    value: 'id'
                },
                {
                    sortable: false,
                    text: '与户主关系',
                    value: 'name'
                },
                {
                    sortable: false,
                    text: '性别',
                    value: 'salary',
                    align: 'right'
                },
                {
                    sortable: false,
                    text: '出生年月',
                    value: 'country',
                    align: 'right'
                },
                {
                    sortable: false,
                    text: '身份证号',
                    value: 'city',
                    align: 'right'
                },
                {
                    sortable: false,
                    text: '手机号',
                    value: 'city',
                    align: 'right'
                },
                {
                    sortable: false,
                    text: '现住址',
                    value: 'city',
                    align: 'right'
                }
            ],
            item: [],
            formData: {
                cityCode: null,
                cityName: null,
                areaCode: null,
                areaName: null,
                villageCode: null,
                villageName: null,
                householder: null,
                contactNumber: null,
                phoneNumber: null,
                familyAddress: null,
                familyType: null,
                residenceBookletCode: null,
                community: null,
                housingType: null,
                housingArea: null,
                wcType: null,
                fuelType: null,
                monthlyPerCapitaIncome: null,
                drinkingWater: null,
                kitchenExhaustType: null,
                airState: null,
                lightingState: null,
                warmAndCold: null,
                healthState: null,
                poultryPen: null,
                founderCode: null,
                founderName: null,
                createUnit: null,
                personLiableCode: null,
                personLiableName: null,
                profileDate: null
            }
        }),
        methods: {
            init () {
                const data = { key: this.$const.FORM_STRUCT.FAMILY_PROFILE_INFO }
                this.$http.get(this.$api.form.struct, data).then(({ data }) => {
                    if (data && data.code === -1) {
                        this.alert({
                            color: 'red',
                            text: '表单结构获取失败！',
                        })
                        return
                    }
                    this.formStruct = data.data.data
                })
            },
            save () {
                const data = this.formData
                this.$http.post(this.$api.familyProfileInfo.save, data).then(({ data }) => {
                    if (data) {
                        this.alert({
                            'text': '嘻嘻嘻嘻'
                        })
                    }
                })
            },
            reset(){
                this.date=null
                this.formData= {
                    cityCode: null,
                      cityName: null,
                      areaCode: null,
                      areaName: null,
                      villageCode: null,
                      villageName: null,
                      householder: null,
                      contactNumber: null,
                      phoneNumber: null,
                      familyAddress: null,
                      familyType: null,
                      residenceBookletCode: null,
                      community: null,
                      housingType: null,
                      housingArea: null,
                      wcType: null,
                      fuelType: null,
                      monthlyPerCapitaIncome: null,
                      drinkingWater: null,
                      kitchenExhaustType: null,
                      airState: null,
                      lightingState: null,
                      warmAndCold: null,
                      healthState: null,
                      poultryPen: null,
                      founderCode: null,
                      founderName: null,
                      createUnit: null,
                      personLiableCode: null,
                      personLiableName: null,
                      profileDate: null
                }
            },
            profileDate (val) {
                this.formData.profileDate = new Date(val).getTime() / 1000
            },
            getRegion (data) {
                this.formData.cityCode = data[0].code
                this.formData.cityName = data[0].name
                this.formData.areaCode = data[1] ? data[1].code : null
                this.formData.areaName = data[1] ? data[1].name : null
                this.formData.villageCode = data[2] ? data[2].code : null
                this.formData.villageName = data[2] ? data[2].name : null
            },
            founderCode (data) {
                this.formData.founderCode = data.value
                this.formData.founderName = data.text
            },
            personLiableCode (data) {
                this.formData.personLiableCode = data.value
                this.formData.personLiableName = data.text
            },
            ...mapMutations([
                'alert'
            ])
        },
        created () {
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
