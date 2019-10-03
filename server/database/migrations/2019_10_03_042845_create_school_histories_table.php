<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateSchoolHistoriesTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('school_histories', function (Blueprint $table) {
            $table->bigIncrements('id');
            $table->unsignedBigInteger('resume_id');
            $table->date('base_date');
            $table->string('detail');
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
        Schema::dropIfExists('school_histories');
    }
}
