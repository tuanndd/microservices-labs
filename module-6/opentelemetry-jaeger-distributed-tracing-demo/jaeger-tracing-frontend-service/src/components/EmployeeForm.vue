<template>
  <div>
    <v-form v-model="valid" id="employee-details-form">
      <v-container>
        <v-row>
          <v-col cols="12" md="4">
            <v-text-field
              v-model="firstName"
              :rules="nameRules"
              :counter="10"
              label="First name"
              required>
            </v-text-field>
          </v-col>
          <v-col cols="12" md="4">
            <v-text-field
              v-model="lastName"
              :rules="nameRules"
              :counter="10"
              label="Last name"
              required>
            </v-text-field>
          </v-col>
          <v-col cols="12" md="4">
            <v-text-field
              v-model="occupation"
              label="Occupation"
              required>
            </v-text-field>
          </v-col>
        </v-row>
      </v-container>
      <v-col class="text-center" cols="18" sm="11">
        <div class="my-2">
          <v-btn 
            form="check-login-form" 
            text 
            large 
            color="primary" 
            @click="submit(); dialog = false" >
            Add Employee
          </v-btn>
          <v-dialog v-model="dialog" width="500">
            <v-card>
              <v-card-title class="headline grey lighten-2">
                <div v-if="isError == true">Error</div>
                <div v-else-if="isError == false">Employee Created</div>
              </v-card-title>
              <v-card-text>
                <div v-if="isError == false">
                  <div text>Employee ID: {{ responseData.id }} </div>
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

      axios.post(NODEJS_SERVICE_URL + 'create-employee/', { 
        firstName: this.firstName,
        lastName: this.lastName,
        occupation: this.occupation
        })
        .then((response) => {
          // handle success
          this.isError = false;
          this.responseData.id = response.data.id;
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
    firstName: '',
    lastName: '',
    nameRules: [
      v => !!v || 'Name is required',
      v => v.length <= 10 || 'Name must be less than 10 characters',
    ],
    occupation: '',
    isError: false,
    errorMessage: '',
    responseData: {}
  }),
}
</script>
