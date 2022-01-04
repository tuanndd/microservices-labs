const express = require('express')
const cors = require('cors')
const axios = require('axios');

const app = express()
app.use(express.json());
app.use(cors());

const port = process.env.SERVER_PORT || 8081;
const JAVA_SERVICE_URL = process.env.JAVA_SERVICE_URL || "http://localhost:8090/";
const GO_SERVICE_URL = process.env.GO_SERVICE_URL || "http://localhost:8091/";
const PYTHON_SERVICE_URL = process.env.PYTHON_SERVICE_URL || "http://localhost:8092/";

async function getGrade(occupation) {
    try {
        const resp = await axios.get(JAVA_SERVICE_URL + 'salary-grade/' + occupation);
        return resp.data.grade;
    } catch (err) {
        throw new Error(err.toString());
    }
}

async function getSalaryAmount(grade) {
    try {
        const resp = await axios.get(PYTHON_SERVICE_URL + 'salary-amount-for-grade/' + grade);
        min = resp.data.minimum;
        max = resp.data.maximum;

        // console.log(min);

        // generate random salary between min and max
        return Math.floor(Math.random() * (max - min + 1)) + min;
    } catch (err) {
        throw new Error(err.toString());
    }
}

async function createEmployee(data) {
    try {
        const resp = await axios.post(GO_SERVICE_URL + 'employee', data);
        return resp.data;
    } catch (err) {
        throw new Error(err.toString());
    } 
}

async function getEmployeeByID(id) {
    try {
        console.log("Getting employee");
        const resp = await axios.get(GO_SERVICE_URL + `employee/${id}`);
        return resp.data;
    } catch (err) {
        throw new Error(err.toString());
    } 
}

app.get('/get-employee/:id', async (req, res) => {
    const id = req.params.id;

    let employee = {};
    try {
        employee = await getEmployeeByID(id);
    } catch (err) {
        console.error(err);
        return res.status(404).send({
            message: 'Employee Not found'
        });
    }

    res.send(employee)
})

app.post('/create-employee', async (req, res) => {
    const body = req.body;
    const occupation = body.occupation;

    let grade = {};
    try {
        grade = await getGrade(occupation);
    } catch (err) {
        console.error(err);
        return res.status(404).send({
            message: 'Not found'
        });
    }

    let salaryAmount = {};
    try {
        salaryAmount = await getSalaryAmount(grade);
    } catch (err) {
        console.error(err);
        return res.status(404).send({
            message: 'Not found'
        });
    }

    let response = {};
    try {
        response = await createEmployee({
            firstName: body.firstName,
            lastName: body.lastName,
            occupation: occupation,
            salaryGrade: grade,
            salaryAmount: salaryAmount.toString(),
        })

    } catch (err) {
        console.error(err);
        return res.status(500).send({
            message: 'Internal Server Error'
        });
    }

    res.send(response)
})

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`)
})
