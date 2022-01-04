<template>
  <div>
    <v-form v-model="valid" id="get-employee-form">
      <v-container>
        <v-row>
          <v-col cols="12" md="4">
            <v-text-field
              v-model="id"
              label="ID"
              required>
            </v-text-field>
          </v-col>
        </v-row>
      </v-container>
      <v-col class="text-center" cols="18" sm="11">
        <div class="my-2">
          <v-btn 
            form="get-employee-form" 
            text 
            large 
            color="primary" 
            @click="submit(); dialog = false" >
            Get Employee
          </v-btn>
          <v-dialog v-model="dialog" width="500">
            <v-card>
              <v-card-title class="headline grey lighten-2">
                <div v-if="isError == true">Error</div>
                <div v-else-if="isError == false">Employee Details</div>
              </v-card-title>
              <v-card-text>
                <div v-if="isError == false">
                  <div text>Employee ID: {{ responseData.id }} </div>
                  <div text>First Name: {{ responseData.firstName }} </div>
                  <div text>Last Name: {{ responseData.lastName }} </div>
                  <div text>Occupation: {{ responseData.occupation }} </div>
                  <div text>Salary Grade: {{ responseData.salaryGrade }} </div>
                  <div text>Salary Amount: {{ responseData.salaryAmount }} </div>
                  <div text>Created At: {{ responseData.createdAt }} </div>
                  <div text>Update At: {{ responseData.updatedAt }} </div>
                </div>
                <div v-else-if="isError == true">
                  <div text>{{ errorMessage }} </div>
                </div>
              </v-card-text>
              <v-divider></v-divider>
              <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="primary" text @click="dialog = false"> OK </v-btn>
              </v-card-actions>
            </v-card>
          </v-dialog>
        </div>
      </v-col>
    </v-form>
  </div>
</template>

<script>
export default {
  methods: {
    submit() {
      const axios = require('axios');

      const NODEJS_SERVICE_URL = process.env.NODEJS_SERVICE_URL || "http://localhost:8081/";

      axios.get(NODEJS_SERVICE_URL + 'get-employee/' + this.id)
        .then((response) => {
          // handle success
          this.isError = false;
          this.responseData.id = response.data.id;
          this.responseData.firstName = response.data.firstName;
          this.responseData.lastName = response.data.lastName;
          this.responseData.occupation = response.data.occupation;
          this.responseData.salaryGrade = response.data.salaryGrade;
          this.responseData.salaryAmount = response.data.salaryAmount;
          this.responseData.createdAt = response.data.created_at;
          this.responseData.updatedAt = response.data.updated_at;
          this.dialog = true;
        })
        .catch((error) => {
          // handle error
          console.log(error);
          this.errorMessage = error;
          this.isError = true;
          this.dialog = true;
        });
    }
  },
  data: () => ({
    dialog: false,
    valid: false,
    isError: false,
    id: '',
    errorMessage: '',
    responseData: {}
  }),
}
</script>
