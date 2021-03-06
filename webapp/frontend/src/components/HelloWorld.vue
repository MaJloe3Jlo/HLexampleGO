<template>
  <v-container fluid>
    <v-snackbar
      top
      v-model="notification.show"
    >
      {{ notification.msg }}
    </v-snackbar>

    <v-row>
      <v-col>
        <v-card class="elevation-12">
          <div class="loader" v-if="card.isLoading">
            <v-progress-circular
              indeterminate
              color="primary"
            ></v-progress-circular>
          </div>
          <v-toolbar color="primary" dark flat>
            <v-toolbar-title>Текущий блок</v-toolbar-title>
          </v-toolbar>
          <v-card-text>
            <v-alert v-if="card.error" type="error" text>
              {{card.error}}
            </v-alert>
            <v-list>
              <v-list-item>
                <v-list-item-icon>
                  <v-icon>mdi-key</v-icon>
                </v-list-item-icon>
                <v-list-item-content>
                  <v-list-item-title>ID: {{ card.model.id || 'пусто' }}</v-list-item-title>
                </v-list-item-content>
              </v-list-item>
              <v-divider></v-divider>
              <v-list-item>
                <v-list-item-icon>
                  <v-icon>mdi-tag-outline</v-icon>
                </v-list-item-icon>
                <v-list-item-content>
                  <v-list-item-title>Название: {{ card.model.name || 'пусто' }}</v-list-item-title>
                </v-list-item-content>
              </v-list-item>
              <v-divider></v-divider>
              <v-list-item>
                <v-list-item-icon>
                  <v-icon>mdi-format-paragraph</v-icon>
                </v-list-item-icon>
                <v-list-item-content>
                  <v-list-item-title>Сумма: {{ card.model.price || 0 }}</v-list-item-title>
                </v-list-item-content>
              </v-list-item>
              <v-divider></v-divider>
              <v-list-item>
                <v-list-item-icon>
                  <v-icon>mdi-zip-box</v-icon>
                </v-list-item-icon>
                <v-list-item-content>
                  <v-list-item-title>Количество {{ card.model.quantity || 0 }}</v-list-item-title>
                </v-list-item-content>
              </v-list-item>
              <v-divider></v-divider>
            </v-list>
          </v-card-text>
          <v-card-actions>
              <v-spacer />
              <v-dialog
                v-model="form.isDialog"
                width="500"
              >
                <template v-slot:activator="{ on, attrs }">
                  <v-btn
                    color="primary"
                    dark
                    v-bind="attrs"
                    v-on="on"
                  >
                    Изменить
                  </v-btn>
                </template>

                <v-card>
                  <div class="loader" v-if="form.isLoading">
                    <v-progress-circular
                      indeterminate
                      color="primary"
                    ></v-progress-circular>
                  </div>

                  <v-card-title class="headline grey lighten-2">
                    Новый блок
                  </v-card-title>

                  <v-card-text class="mt-5">
                    <v-alert v-if="form.error" type="error" text>
                      {{form.error}}
                    </v-alert>
                    <v-form
                      ref="form"
                      v-model="form.valid"
                    >
                      <v-text-field
                        v-model="form.model.id"
                        label="ID"
                        name="id"
                        prepend-icon="mdi-key"
                        type="number"
                        :rules="validator.required"
                        required
                      />                    
                      <v-text-field
                        v-model="form.model.name"
                        label="Название"
                        name="name"
                        prepend-icon="mdi-tag-outline"
                        type="text"
                        :rules="validator.required"
                        required
                      />
                      <v-text-field
                        v-model="form.model.price"
                        label="Сумма"
                        name="price"
                        prepend-icon="mdi-format-paragraph"
                        type="number"
                        :rules="validator.required"
                        required
                      />
                      <v-text-field
                        v-model="form.model.quantity"
                        label="Количество"
                        name="quantity"
                        prepend-icon="mdi-zip-box"
                        type="number"
                        :rules="validator.required"
                        required
                      />
                    </v-form>
                  </v-card-text>

                  <v-divider></v-divider>

                  <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn
                      color="primary"
                      @click="onSubmit"
                    >
                      Сохранить
                    </v-btn>
                  </v-card-actions>
                </v-card>
              </v-dialog>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <v-card class="elevation-12">
          <div class="loader" v-if="list.isLoading">
            <v-progress-circular
              indeterminate
              color="primary"
            ></v-progress-circular>
          </div>
          <v-toolbar color="primary" dark flat>
            <v-toolbar-title>Список блоков</v-toolbar-title>
          </v-toolbar>
          <v-card-text>
            <v-alert v-if="list.error" type="error" text>
              {{list.error}}
            </v-alert>
            <v-simple-table v-if="list.data">
              <template v-slot:default>
                <thead>
                  <tr>
                    <th class="text-left">
                      id
                    </th>
                    <th class="text-left">
                      Название
                    </th>
                    <th class="text-left">
                      Сумма
                    </th>
                    <th class="text-left">
                      Количество
                    </th>
                  </tr>
                </thead>
                <tbody>
                  <tr
                    v-for="item in list.data"
                    :key="item.id"
                  >
                    <td>{{ item.id }}</td>
                    <td>{{ item.name }}</td>
                    <td>{{ item.price }}</td>
                    <td>{{ item.quantity }}</td>
                  </tr>
                </tbody>
              </template>
            </v-simple-table>
            <v-alert v-else class="mb-0" type="warning" text>
              Нет данных
            </v-alert>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
  import api from '@/api';

  const MODEL = {
    id: '',
    name: '',
    price: '',
    quantity: ''
  };

  export default {
    name: 'HelloWorld',

    data() {
      return {
        card: {
          model: this.setModel(),
          error: '',
          isLoading: false,
        },
        form: {
          isLoading: false,
          isDialog: false,
          error: '',
          valid: false,
          model: this.setModel()
        },
        list: {
          data: null,
          isLoading: false,
          error: ''
        },
        validator: {
          required: [
            v => !!v || 'Обязательное поле'
          ]
        },
        notification: {
          msg: 'Готово',
          show: false
        }
      }
    },

    methods: {
      setModel(model) {
        return Object.assign({}, model || MODEL)
      },

      find() {
        this.list.error = '';
        this.list.isLoading = true;
        api.list().then(({data}) => {
          if (data) {
            this.list.data = data;
          } else {
            this.list.error = 'Нет данных'
          }
        }, err => {
          this.card.error = err || 'Ошибка';
        }).finally(() => {
         this.list.isLoading = false;
        });
      },

      findOne() {
        this.card.error = '';
        this.card.isLoading = true;
        api.get().then(({data}) => {
          if (data) {
            this.card.model = this.setModel(data);
          } else {
            this.card.error = 'Нет данных'
          }
        }, err => {
          this.card.error = err || 'Ошибка';
        }).finally(() => {
         this.card.isLoading = false;
        });
      },

      onSubmit() {
        this.form.error = '';
        this.$refs.form.validate()
        if (!this.form.valid) return false;

        this.form.isLoading = true;

        const { id, name, price, quantity } = this.form.model;

        api.save({id: +id, name, price: +price, quantity: +quantity}).then(({data}) => {
          this.form.isDialog = false;
          this.form.model = this.setModel();
          this.notification.msg = 'ID транзации: ' + data;
          this.notification.show = true;
          this.findOne();
          this.find();
        }, err => {
          this.form.error = err || 'Ошибка';
        }).finally(() => {
          this.form.isLoading = false;
        });
      }
    },

    created() {
      this.findOne();
      this.find();
    }
  }
</script>

<style scoped>
  .loader {
    position: absolute;
    width: 100%;
    height: 100%;
    z-index: 999;
    background: rgba(255, 255, 255, 0.87);
    display: flex;
    align-items: center;
    justify-content: center;
    top: 0;
    left: 0;
  }
  .container {
    max-width: 1200px;
  }
</style>