<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateResumesTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('resumes', function (Blueprint $table) {
            $table->bigIncrements('id');
            $table->unsignedBigInteger('user_id')->nullable(false);
            $table->string('kana');
            $table->binary('photo');
            $table->date('birthday');
            $table->string('house_tel');
            $table->string('cellular_phone');
            $table->string('fax');
            $table->string('contact_email');
            $table->string('zip_code');
            $table->unsignedBigInteger('prefecture');
            $table->unsignedBigInteger('municipality');
            $table->string('address_1');
            $table->string('address_2');
            $table->multiLineString('description');
            $table->string('created_by');
            $table->string('updated_by');
            $table->timestamps();
            $table->softDeletes();
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('resumes');
    }
}
